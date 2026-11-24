# Maximum XOR of Two Numbers

> Striver SDE Sheet 190 · Day 27 Trie — a binary trie over the bits; for each number, walk it greedily taking the opposite bit whenever one exists. Go, stdlib only, tested.

## Problem

The largest XOR obtainable from any pair in an array — O(n²) by brute force; the trie makes it O(n) in word-sized bits.

## Approach

Store numbers as 32-level paths in a 2-ary trie, most significant bit first. To maximize `v ^ ?`, greed is provably right bit by bit: a 1 in a higher position beats *any* setting of all lower positions, so at each level take the child with the **opposite** bit when it exists. Each query is 32 steps regardless of n. Insert-then-query-as-you-go covers all pairs once. This binary-trie-with-greedy-descent is a mini-paradigm — it reappears whenever "maximize/minimize XOR against a set" shows up, including the very next problem. A brute-force cross-check anchors the tests.

**Complexity:** O(n · 32) time, O(n · 32) trie nodes.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
