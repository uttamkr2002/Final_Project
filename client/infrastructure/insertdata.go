package infrastructure

import (
	models "client/modelsWithInterface"
	"fmt"
)

// Improved version
func InsertMetrics(dbClient SqlServices, payload models.Payload) (int, error) {
	// Check connection first
	err := dbClient.Ping()
	if err != nil {
		return 0, fmt.Errorf("database connection error: %v", err)
	}

	fmt.Println("📝 Inserting collected metrics into PostgreSQL...")

	// Define the SQL INSERT statement with placeholders and RETURNING id
	query := `
        INSERT INTO mydata (
            name, disk_total, disk_used, iops_in_progress, 
            swap_total, swap_used, virtual_total, virtual_used, 
            buffers, cached, uptime, platform, platform_version, cpu_usage
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
        )
        RETURNING id
    `

	// Execute the query with collected metrics
	var id int
	row := dbClient.QueryRow(query,
		"Server1",
		payload.Disk.Total,          // Disk Total
		payload.Disk.Used,           // Disk Used
		payload.Disk.IopsInProgress, // IOPS
		payload.Memory.SwapTotal,    // Swap Total
		payload.Memory.SwapUsed,     // Swap Used
		payload.Memory.VirtualTotal, // Virtual Total
		payload.Memory.VirtualUsed,  // Virtual Used
		payload.Memory.Buffers,      // Buffers
		payload.Memory.Cached,       // Cached
		payload.OS.Uptime,           // Uptime
		payload.OS.Platform,         // Platform
		payload.OS.PlatformVersion,  // Platform Version
		payload.CPU.CPUUsage,        // CPU Usage
	)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting metrics: %v", err)
	}

	fmt.Println("✅ Metrics inserted successfully!")
	return id, nil
}
