package worker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPool_SubmitAndWait(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool := NewPool(2)
	pool.Start(ctx)

	// ジョブを投入
	for i := 0; i < 5; i++ {
		job := &Job{
			ID:   fmt.Sprintf("JOB-%d", i),
			Data: fmt.Sprintf("data-%d", i),
		}
		if err := pool.Submit(ctx, job); err != nil {
			t.Fatalf("ジョブの投入に失敗: %v", err)
		}
	}

	pool.Wait()

	metrics := pool.GetMetrics()
	if metrics.SuccessCount+metrics.FailureCount != 5 {
		t.Errorf("処理されたジョブ数が期待と異なります: 期待=5, 実際=%d", metrics.SuccessCount+metrics.FailureCount)
	}
}

func TestPool_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	pool := NewPool(2)
	pool.Start(ctx)

	// ジョブを投入
	job := &Job{ID: "JOB-1", Data: "data-1"}
	pool.Submit(ctx, job)

	// すぐにキャンセル
	cancel()

	// 少し待ってから終了
	time.Sleep(100 * time.Millisecond)
	pool.Wait()
}

