package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"worker_pool/internal/worker"
)

func main() {
	numWorkers := flag.Int("workers", 3, "ワーカーの数")
	numJobs := flag.Int("jobs", 10, "ジョブの数")
	timeout := flag.Duration("timeout", 30*time.Second, "タイムアウト時間")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	pool := worker.NewPool(*numWorkers)
	pool.Start(ctx)

	// ジョブを投入
	for i := 0; i < *numJobs; i++ {
		jobID := fmt.Sprintf("JOB-%03d", i+1)
		pool.Submit(ctx, &worker.Job{
			ID:   jobID,
			Data: fmt.Sprintf("data-%d", i+1),
		})
	}

	// すべてのジョブが完了するまで待機
	pool.Wait()

	// 結果を表示
	metrics := pool.GetMetrics()
	log.Printf("処理完了: 成功=%d, 失敗=%d, 合計時間=%v",
		metrics.SuccessCount, metrics.FailureCount, metrics.TotalDuration)
}

