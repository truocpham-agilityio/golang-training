/*
You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth.

The coins will be distributed based on the vowels contained in each name where:

a: 1 coin

e: 1 coin

i: 2 coins

o: 3 coins

u: 4 coins

A user can’t get more than 10 coins. For example, for ‘Augutus’ the total count of coins is 13. But he’ll get only 10 (13-3) coins.

Print a map with each user’s name and the number of coins distributed. After distributing all the coins, you should have 2 coins left.

The output should look something like that:

map[Matthew:2 Peter:2 Giana:4 Adriano:7 Elizabeth:5 Sarah:2 Augustus:10 Heidi:5 Emilie:6 Aaron:5]
Coins left: 2
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func GetResult([] string) string {
	coinHash := make(map[string]int, 5)
	coinHash["a"] = 1
	coinHash["e"] = 1
	coinHash["i"] = 2
	coinHash["o"] = 3
	coinHash["u"] = 4

	for _, user := range users {
		count := 0
		for _, char := range user {
			if _, ok := coinHash[strings.ToLower(string(char))]; ok {
				count += coinHash[strings.ToLower(string(char))]
			}
		}
		if count > 10 {
			count = 10
		}
		distribution[user] = count
		coins = coins - count
	}

  	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
  	
	return strconv.Itoa(distribution["Matthew"])+" "+strconv.Itoa(distribution["Sarah"])+" "+strconv.Itoa(distribution["Augustus"])+" "+strconv.Itoa(distribution["Heidi"])+" "+strconv.Itoa(distribution["Emilie"])+" "+strconv.Itoa(distribution["Peter"])+" "+strconv.Itoa(distribution["Giana"])+" "+strconv.Itoa(distribution["Adriano"])+" "+strconv.Itoa(distribution["Aaron"])+" "+strconv.Itoa(distribution["Elizabeth"])
}
