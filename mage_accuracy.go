//attempt to calculate accuracy in pvp combat

package main

import (
	"fmt"
	"math"
	"strconv"
)

type Prayer struct {
	name    string
	magemul float64
	defmul  float64
}

var prayers [6]Prayer

type Player struct {
	num            int
	l_def          int
	l_mage         int
	magic_attack   int
	magic_defense  int
	magic_prayer   float64
	defense_prayer float64
}

func prayer_array_init() {
	prayers[0] = Prayer{"Augury", 1.25, 1.25}
	prayers[1] = Prayer{"Piety", 1.00, 1.25}
	prayers[2] = Prayer{"Mystic Might + Steel Skin", 1.15, 1.15}
	prayers[3] = Prayer{"Mystic Might", 1.15, 1.00}
	prayers[4] = Prayer{"Steel Skin", 1.00, 1.15}
	prayers[5] = Prayer{"No prayers", 1.00, 1.00}
}

func print_prayers() {
	i := 0
	for i < 6 {
		fmt.Println(strconv.Itoa(i) + ": " + prayers[i].name)
		i += 1
	}
}
func player_set_levels(player *Player, pname string) {
	var deflvl, magelvl int
	setprint := false
	fmt.Printf("select defense level of " + pname + ": ")
	fmt.Scanf("%d\n", &deflvl)
	if deflvl < 1 {
		deflvl = 1
		setprint = true
	}
	if setprint {
		fmt.Println("invalid defence lvl, set to " + strconv.Itoa(deflvl))
	}
	fmt.Printf("select magic level of " + pname + ": ")
	fmt.Scanf("%d\n", &magelvl)
	setprint = false
	if magelvl < 1 {
		magelvl = 1
		setprint = true
	}
	if setprint {
		fmt.Println("invalid magic level, set to " + strconv.Itoa(magelvl))
	}
	player.l_def = deflvl
	player.l_mage = magelvl
}
func player_set_prayers(player *Player, pname string) {
	setprint := false
	var prayerindex int
	print_prayers()
	fmt.Println("select player" + pname + "'s prayers")
	fmt.Scanf("%d\n", &prayerindex)
	if prayerindex > 4 {
		prayerindex = 4
		setprint = true
	} else if prayerindex < 0 {
		prayerindex = 0
		setprint = true
	}
	if setprint {
		fmt.Println("invalid prayerindex, set to " + strconv.Itoa(prayerindex) + ", " + prayers[prayerindex].name)
	}
	player.magic_prayer = prayers[prayerindex].magemul
	player.defense_prayer = prayers[prayerindex].defmul
}

func player_set_magic_stats(player *Player, pname string) {
	var magic_attack, magic_defense int
	fmt.Printf("magic attack of " + pname + ": ")
	fmt.Scanf("%d\n", &magic_attack)
	fmt.Printf("magic defense of " + pname + ": ")
	fmt.Scanf("%d\n", &magic_defense)
	player.magic_attack = magic_attack
	player.magic_defense = magic_defense
}

func player_setup(player *Player) {
	pname := "player" + strconv.Itoa(player.num)
	player_set_levels(player, pname)
	player_set_prayers(player, pname)
	player_set_magic_stats(player, pname)
	fmt.Println(strconv.Itoa(player.l_def) + " " + strconv.Itoa(player.l_mage) + " " + strconv.Itoa(player.magic_attack) + " " + strconv.Itoa(player.magic_defense))
}

func effective_level(level int, mul float64) int {
	eff_level := int(math.Floor(float64(level) * mul))
	//no accuracy bonuses from stance
	eff_level += 8
	//no void magic bonus implemented
	return eff_level
}

func magic_attack_roll_max(player *Player) int {
	A := effective_level(player.l_mage, player.magic_prayer)
	max_roll := A * (player.magic_attack + 64)
	return max_roll
}

func magic_defense_roll_max(player *Player) int {
	eff_defense_level := int(math.Floor(0.30 * float64(effective_level(player.l_def, player.defense_prayer))))
	eff_magic_level := int(math.Floor(0.70 * float64(math.Floor(float64(player.l_mage)*player.magic_prayer))))
	A := eff_defense_level + eff_magic_level
	B := player.magic_defense
	max_roll := A * (B + 64)
	return max_roll
}

func hit_chance(attacker *Player, defender *Player) float64 {
	var hit_chance float64
	attack_roll := float64(magic_attack_roll_max(attacker))
	defense_roll := float64(magic_defense_roll_max(defender))
	if attack_roll > defense_roll {
		hit_chance = 1.0 - (defense_roll+2.0)/(2.0*(attack_roll+1.0))
	} else {
		hit_chance = (attack_roll / (2.0 * (defense_roll + 1.0)))
	}
	return hit_chance
}

func main() {
	prayer_array_init()
	p0 := new(Player)
	p1 := new(Player)
	//initialize each player's magic stats
	p0.num = 0
	p1.num = 1
	player_setup(p0)
	player_setup(p1)
	hit_chance_0 := hit_chance(p0, p1)
	hit_chance_1 := hit_chance(p1, p0)
	fmt.Printf("player 0 has a %f chance of hitting player 1\n", hit_chance_0)
	fmt.Printf("player 1 has a %f chance of hitting player 0\n", hit_chance_1)

}
