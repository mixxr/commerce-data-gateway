package models

import (
	"bytes"
	"fmt"
)

const VALUES_SUFFIX string = "values"

type TableValues struct {
	table_ *Table
	Start  int
	Count  int
	Rows   [][]string
}

func NewValues(t *Table, start int, count int, rows [][]string) *TableValues {
	newObj := new(TableValues)
	newObj.table_ = t
	newObj.Start = start
	newObj.Count = count
	if rows != nil {
		newObj.Rows = rows
		newObj.Count = len(rows)
	}
	return newObj
}

func (o *TableValues) Size() int {
	return len(o.Rows) * len(o.Rows[0])
}

func (o *TableValues) String() string {
	return fmt.Sprintf("%d,%d", o.Start, o.Count)
}

func (o *TableValues) Parent() *Table {
	return o.table_
}

// returns SQL instruction: CREATE TABLE <owner>_<name>_values ...
// example
// CREATE TABLE mike_ssn_ca_values (
// 	id BIGINT NOT NULL AUTO_INCREMENT,
// 	created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	colvalue1 VARCHAR(256) NOT NULL,
// 	colvalue2 VARCHAR(256) NOT NULL,
// 	colvalue3 VARCHAR(256) NOT NULL,
// 	PRIMARY KEY ( id )
//  );
func (o *TableValues) GetCreateTable() (string, error) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, "CREATE TABLE %s_%s_%s (", o.table_.Owner, o.table_.Name, VALUES_SUFFIX)

	buffer.WriteString("id BIGINT NOT NULL AUTO_INCREMENT,")
	buffer.WriteString("created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,")

	for k := 0; k < len(o.Rows[0]); k++ {
		fmt.Fprintf(&buffer, "colvalue%d VARCHAR(256) NOT NULL,", k)
	}

	buffer.WriteString("PRIMARY KEY ( id ));")

	return buffer.String(), nil
}

// returns SQL instruction: INSERT INTO  <owner>_<name>_values ...
// Example:
// INSERT INTO mike_ssn_ca (colvalue1,colvalue2,colvalue3) VALUES
// ('mike','douglàs','3897428934EWREW'),
// ('äbel','òmar ópël','3897428934EWREW'),
// ('anthony','di martino','234234FSAFSADF');

func (o *TableValues) GetInsertTable() (string, error) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, "INSERT INTO %s_%s_%s (", o.table_.Owner, o.table_.Name, VALUES_SUFFIX)
	i, j := 0, 0
	for ; i < len(o.Rows[0])-1; i++ {
		fmt.Fprintf(&buffer, "colvalue%d,", i)
	}
	fmt.Fprintf(&buffer, "colvalue%d) VALUES ", i)
	i = 0
	for ; i < len(o.Rows)-1; i++ {
		buffer.WriteString("(")
		j = 0
		for ; j < len(o.Rows[i])-1; j++ {
			fmt.Fprintf(&buffer, "'%s',", o.Rows[i][j])
		}
		fmt.Fprintf(&buffer, "'%s'),", o.Rows[i][j])
	}
	buffer.WriteString("(")
	j = 0
	for ; j < len(o.Rows[i])-1; j++ {
		fmt.Fprintf(&buffer, "'%s',", o.Rows[i][j])
	}
	fmt.Fprintf(&buffer, "'%s');", o.Rows[i][j])

	return buffer.String(), nil
}