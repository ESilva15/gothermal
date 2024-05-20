-- Generate a password
-- htpasswd -bnBC 10 "" kika | tr -d ':\n'

CREATE TABLE IF NOT EXISTS users(
	id       INTEGER PRIMARY KEY AUTOINCREMENT,
	name     TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
);

-- pass is kika here
INSERT INTO
  users(name, password)
	values(
		'kika',
		'$2y$10$JgmGJBRPqJB5/S9gL24nC.JuMj.wnFmUHBwENCGFboCzeNATmu.rq%'
		);