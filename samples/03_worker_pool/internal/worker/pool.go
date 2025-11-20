package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Job は処理対象のジョブです
type Job struct {
	ID   string
	Data string
}

// Result はジョブの処理結果です
type Result struct {
	JobID     string
	Success   bool
	Error     error
	Duration  time.Duration
	ProcessedAt time.Time
}

// Metrics は処理メトリクスです
type Metrics struct {
	SuccessCount int
	FailureCount int
	TotalDuration time.Duration
}

// Pool はワーカープールです
type Pool struct {
	workers    int
	jobQueue   chan *Job
	resultChan chan *Result
	wg         sync.WaitGroup
	metrics    *Metrics
	mu         sync.Mutex
	startTime  time.Time
}

// NewPool は新しいワーカープールを作成します
func NewPool(workers int) *Pool {
	return &Pool{
		workers:    workers,
		jobQueue:   make(chan *Job, workers*2),
		resultChan: make(chan *Result, workers*2),
		metrics:    &Metrics{},
	}
}

// Start はワーカープールを開始します
func (p *Pool) Start(ctx context.Context) {
	p.startTime = time.Now()

	// ワーカーを起動
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}

	// 結果収集goroutine
	p.wg.Add(1)
	go p.collectResults(ctx)
}

// Submit はジョブを投入します
func (p *Pool) Submit(ctx context.Context, job *Job) error {
	select {
	case p.jobQueue <- job:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Wait はすべてのジョブが完了するまで待機します
func (p *Pool) Wait() {
	close(p.jobQueue)
	p.wg.Wait()
	close(p.resultChan)
}

// GetMetrics はメトリクスを取得します
func (p *Pool) GetMetrics() *Metrics {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.metrics.TotalDuration = time.Since(p.startTime)
	return p.metrics
}

func (p *Pool) worker(ctx context.Context, id int) {
	defer p.wg.Done()

	for {
		select {
		case job, ok := <-p.jobQueue:
			if !ok {
				return
			}
			p.processJob(ctx, job, id)
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pool) processJob(ctx context.Context, job *Job, workerID int) {
	start := time.Now()

	// ジョブ処理をシミュレート
	result := &Result{
		JobID:       job.ID,
		ProcessedAt: time.Now(),
	}

	// 実際の処理（ここではシミュレーション）
	err := p.executeJob(ctx, job, workerID)
	if err != nil {
		result.Success = false
		result.Error = err
	} else {
		result.Success = true
	}

	result.Duration = time.Since(start)

	select {
	case p.resultChan <- result:
	case <-ctx.Done():
		return
	}
}

func (p *Pool) executeJob(ctx context.Context, job *Job, workerID int) error {
	// 処理をシミュレート（ランダムな遅延）
	select {
	case <-time.After(time.Duration(100+workerID*50) * time.Millisecond):
		// 10%の確率でエラーを発生
		if time.Now().UnixNano()%10 == 0 {
			return fmt.Errorf("処理エラー: %s", job.ID)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (p *Pool) collectResults(ctx context.Context) {
	defer p.wg.Done()

	for {
		select {
		case result, ok := <-p.resultChan:
			if !ok {
				return
			}
			p.updateMetrics(result)
			if result.Success {
				fmt.Printf("[成功] %s (処理時間: %v)\n", result.JobID, result.Duration)
			} else {
				fmt.Printf("[失敗] %s: %v (処理時間: %v)\n", result.JobID, result.Error, result.Duration)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pool) updateMetrics(result *Result) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if result.Success {
		p.metrics.SuccessCount++
	} else {
		p.metrics.FailureCount++
	}
}

