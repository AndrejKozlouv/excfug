	// Create a new shuffler instance.
	shuffler := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Shuffle the list of guilds.
	shuffler.Shuffle(len(list_guilds), func(i, j int) {
		list_guilds[i], list_guilds[j] = list_guilds[j], list_guilds[i]
	})  
