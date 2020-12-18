package snapshotArray

type SnapshotArray struct {
	snapId int
	m      map[int]int
	mslice []map[int]int
}

func Constructor(length int) SnapshotArray {
	mslice := []map[int]int{}
	m := make(map[int]int)
	return SnapshotArray{0, m, mslice}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.m[index] = val
}

func (this *SnapshotArray) Snap() int {
	this.mslice = append(this.mslice, this.m)
	this.m = make(map[int]int)
	snapId := this.snapId
	this.snapId++
	return snapId
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	val := 0
	for i := 0; i <= snap_id; i++ {
		if v, ok := this.mslice[i][index]; ok {
			val = v
		}
	}

	return val
}
