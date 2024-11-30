package src

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HandleCurrentTime handles the /current-time endpoint
func HandleCurrentTime(c *gin.Context, db *sql.DB) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load time zone"})
		return
	}
	currentTime := time.Now().In(loc)

	// Insert time into the database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log time"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"current_time": currentTime.Format(time.RFC3339)})
}

// HandleAllTimes handles the /all-times endpoint
func HandleAllTimes(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT timestamp FROM time_log")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve times"})
		return
	}
	defer rows.Close()

	var times []string
	for rows.Next() {
		var ts time.Time
		if err := rows.Scan(&ts); err != nil {
			continue
		}
		times = append(times, ts.Format(time.RFC3339))
	}

	c.JSON(http.StatusOK, gin.H{"times": times})
}
