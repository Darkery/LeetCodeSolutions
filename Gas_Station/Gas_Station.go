package Gas_Station

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}
		total := gas[i] - cost[i]
		j := i + 1
		for j != i{
			if j >= len(gas) {
				j = 0
			}
			total = total + gas[j] - cost[j]
			if total < 0 {
				break
			}
			if j == i {
				return i
			}
			j++
		}
		if j == i {
			return i
		}
	}

	return -1
}
