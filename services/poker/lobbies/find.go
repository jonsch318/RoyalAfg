package lobbies

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"

func find(class []string, lobbies map[string]*lobby.Lobby) string {
	options := make([]struct {
		id    string
		count int
	}, 0)

	for _, n := range class {
		l, ok := lobbies[n]
		if ok {
			count := l.TotalPlayerCount()
			if count < 9 {
				j := 0
				for i := range options {
					if options[i].count > count {
						j = i
						break
					}
				}

				if len(options) > j+1 {
					options = append(options[:j+1], options[j:]...)
					options[j] = struct {
						id    string
						count int
					}{id: n, count: count}
				} else {
					options = append(options, struct {
						id    string
						count int
					}{id: n, count: count})
				}

			}
		}
	}

	if len(options) > 0 {
		return options[0].id
	}

	return ""

}
