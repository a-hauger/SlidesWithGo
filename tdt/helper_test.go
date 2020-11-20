package tdt

import "testing"

func TestDatabase(t *testing.T) {
	//some setup code
	for _, tc := range testCases {
		t.Run(tc.name, tc.DB.testDatabase_Find(t))
		t.Run(tc.name, func(t *testing.T) {
			//more test code
		})
	}
}
func (d *psql) testDatabase_find(t *testing.T) {
	//code
}
