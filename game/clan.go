package game

type Clan struct {
	BadgeUrls struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"badgeUrls"`
	ClanLevel        int64  `json:"clanLevel"`
	ClanPoints       int64  `json:"clanPoints"`
	ClanVersusPoints int64  `json:"clanVersusPoints"`
	Description      string `json:"description"`
	IsWarLogPublic   bool   `json:"isWarLogPublic"`
	Location         struct {
		CountryCode string `json:"countryCode"`
		ID          int64  `json:"id"`
		IsCountry   bool   `json:"isCountry"`
		Name        string `json:"name"`
	} `json:"location"`
	Members                int64  `json:"members"`
	Name                   string `json:"name"`
	RequiredTownhallLevel  int64  `json:"requiredTownhallLevel"`
	RequiredTrophies       int64  `json:"requiredTrophies"`
	RequiredVersusTrophies int64  `json:"requiredVersusTrophies"`
	Tag                    string `json:"tag"`
	Type                   string `json:"type"`
	WarFrequency           string `json:"warFrequency"`
	WarLeague              struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"warLeague"`
	WarLosses    int64 `json:"warLosses"`
	WarTies      int64 `json:"warTies"`
	WarWinStreak int64 `json:"warWinStreak"`
	WarWins      int64 `json:"warWins"`
	MemberList   []ClanMember
}

type ClanMember struct {
	ClanRank          int64 `json:"clanRank"`
	Donations         int64 `json:"donations"`
	DonationsReceived int64 `json:"donationsReceived"`
	ExpLevel          int64 `json:"expLevel"`
	League            struct {
		IconUrls struct {
			Medium string `json:"medium"`
			Small  string `json:"small"`
			Tiny   string `json:"tiny"`
		} `json:"iconUrls"`
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"league"`
	Name             string `json:"name"`
	PreviousClanRank int64  `json:"previousClanRank"`
	Role             string `json:"role"`
	Tag              string `json:"tag"`
	Trophies         int64  `json:"trophies"`
	VersusTrophies   int64  `json:"versusTrophies"`
}

type Member struct {
	Achievements []struct {
		CompletionInfo string `json:"completionInfo"`
		Info           string `json:"info"`
		Name           string `json:"name"`
		Stars          int64  `json:"stars"`
		Target         int64  `json:"target"`
		Value          int64  `json:"value"`
		Village        string `json:"village"`
	} `json:"achievements"`
	AttackWins         int64 `json:"attackWins"`
	BestTrophies       int64 `json:"bestTrophies"`
	BestVersusTrophies int64 `json:"bestVersusTrophies"`
	BuilderHallLevel   int64 `json:"builderHallLevel"`
	Clan               struct {
		BadgeUrls struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Small  string `json:"small"`
		} `json:"badgeUrls"`
		ClanLevel int64  `json:"clanLevel"`
		Name      string `json:"name"`
		Tag       string `json:"tag"`
	} `json:"clan"`
	DefenseWins       int64 `json:"defenseWins"`
	Donations         int64 `json:"donations"`
	DonationsReceived int64 `json:"donationsReceived"`
	ExpLevel          int64 `json:"expLevel"`
	Heroes            []struct {
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Name     string `json:"name"`
		Village  string `json:"village"`
	} `json:"heroes"`
	Labels []interface{} `json:"labels"`
	League struct {
		IconUrls struct {
			Medium string `json:"medium"`
			Small  string `json:"small"`
			Tiny   string `json:"tiny"`
		} `json:"iconUrls"`
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"league"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Spells []struct {
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Name     string `json:"name"`
		Village  string `json:"village"`
	} `json:"spells"`
	Tag           string `json:"tag"`
	TownHallLevel int64  `json:"townHallLevel"`
	Troops        []struct {
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Name     string `json:"name"`
		Village  string `json:"village"`
	} `json:"troops"`
	Trophies             int64 `json:"trophies"`
	VersusBattleWinCount int64 `json:"versusBattleWinCount"`
	VersusBattleWins     int64 `json:"versusBattleWins"`
	VersusTrophies       int64 `json:"versusTrophies"`
	WarStars             int64 `json:"warStars"`
}

type Result struct {
	data []Member
}
