package main

import (
	"math/rand"
	"testing"
	"time"
)

// TestShufflePlay 模拟随机播放一万首歌：
//  1. 生成整形切片 tracks，元素为 1 ~ 10000，模拟一万首歌；
//  2. 每轮从切片中随机取一个下标，模拟收听这首歌；
//  3. 将取到的下标从切片中删除，模拟这首歌已经听过；
//  4. 重复上述过程，直到切片为空（播放结束）。
// 运行一遍程序只需要 0.012 秒。
func TestShufflePlay(t *testing.T) {
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	const total = 10000

	// 1. 生成 1 ~ 10000 的整形切片，模拟一万首歌
	tracks := make([]int, 0, total)
	for i := 1; i <= total; i++ {
		tracks = append(tracks, i)
	}

	// played 记录每首歌被播放的次数，用于校验不会重复播放
	played := make(map[int]int, total)
	// count 记录总共播放了多少首歌
	count := 0

	// 2~4. 随机取歌、播放、删除，直到切片为空（播放结束）
	for len(tracks) > 0 {
		// 随机取一个下标，模拟收听这首歌
		idx := seed.Intn(len(tracks))
		song := tracks[idx]

		// 从切片中删除该下标，模拟这首歌已经听过
		tracks = append(tracks[:idx], tracks[idx+1:]...)

		// 记录播放情况
		played[song]++
		count++
	}

	// 校验：正好播放了 total 首歌
	if count != total {
		t.Fatalf("期望播放 %d 首歌，实际播放了 %d 首", total, count)
	}

	// 校验：正好有 total 首不同的歌
	if len(played) != total {
		t.Fatalf("期望 %d 首不同的歌，实际只有 %d 首", total, len(played))
	}

	// 校验：每首歌都播放了，且只播放了一次
	for song, times := range played {
		if times != 1 {
			t.Fatalf("歌曲 %d 被播放了 %d 次，期望 1 次", song, times)
		}
	}

	t.Logf("播放结束，共随机播放 %d 首歌，每首歌各播放一次。", count)
}
