package Q100

func Quest() string {
	return `select top(1) with ties player_id, event_date as first_login 
			from activity 
			order by row_number() over(partition by player_id order by event_date);`
}

func Quest2() string {
	return `select player_id, min(event_date) as first_login
			from Activity
			group by player_id;`
}
