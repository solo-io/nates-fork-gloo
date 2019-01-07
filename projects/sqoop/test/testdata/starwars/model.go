package starwars

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CharacterFields struct {
	TypeName  string    `json:"__typename"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	FriendIds []string  `json:"friend_ids"`
	AppearsIn []Episode `json:"appears_in"`
}

type Human struct {
	CharacterFields
	StarshipIds  []string `json:"starship_ids"`
	heightMeters float64  `json:"height_meters"`
	Mass         float64  `json:"mass"`
}

func (h *Human) Height(unit LengthUnit) float64 {
	switch unit {
	case "METER", "":
		return h.heightMeters
	case "FOOT":
		return h.heightMeters * 3.28084
	default:
		panic("invalid unit")
	}
}

type Starship struct {
	ID           string
	Name         string
	History      [][]int
	lengthMeters float64
}

func (s *Starship) Length(unit LengthUnit) float64 {
	switch unit {
	case "METER", "":
		return s.lengthMeters
	case "FOOT":
		return s.lengthMeters * 3.28084
	default:
		panic("invalid unit")
	}
}

type Review struct {
	Stars      int
	Commentary *string
	Time       time.Time
}

type Droid struct {
	CharacterFields
	PrimaryFunction string
}

func (r *Resolver) resolveFriendConnection(ctx context.Context, ids []string, first *int, after *string) (FriendsConnection, error) {
	from := 0
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return FriendsConnection{}, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return FriendsConnection{}, err
		}
		from = i
	}

	to := len(ids)
	if first != nil {
		to = from + *first
		if to > len(ids) {
			to = len(ids)
		}
	}

	return FriendsConnection{
		ids:  ids,
		from: from,
		to:   to,
	}, nil
}

type FriendsConnection struct {
	ids  []string
	from int
	to   int
}

func (f *FriendsConnection) TotalCount() int {
	return len(f.ids)
}

func (f *FriendsConnection) PageInfo() PageInfo {
	return PageInfo{
		StartCursor: encodeCursor(f.from),
		EndCursor:   encodeCursor(f.to - 1),
		HasNextPage: f.to < len(f.ids),
	}
}

func encodeCursor(i int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1)))
}

type FriendsEdge struct {
	Cursor string
	Node   Character
}
