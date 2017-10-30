package acca

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type invoiceTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("finances").
func (v *invoiceTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("invoices").
func (v *invoiceTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *invoiceTableType) Columns() []string {
	return []string{"invoice_id", "order_id", "destination_id", "source_id", "paid", "amount", "created_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *invoiceTableType) NewStruct() reform.Struct {
	return new(Invoice)
}

// NewRecord makes a new record for that table.
func (v *invoiceTableType) NewRecord() reform.Record {
	return new(Invoice)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *invoiceTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// InvoiceTable represents invoices view or table in SQL database.
var InvoiceTable = &invoiceTableType{
	s: parse.StructInfo{Type: "Invoice", SQLSchema: "finances", SQLName: "invoices", Fields: []parse.FieldInfo{{Name: "InvoiceID", PKType: "int64", Column: "invoice_id"}, {Name: "OrderID", PKType: "", Column: "order_id"}, {Name: "DestinationID", PKType: "", Column: "destination_id"}, {Name: "SourceID", PKType: "", Column: "source_id"}, {Name: "Paid", PKType: "", Column: "paid"}, {Name: "Amount", PKType: "", Column: "amount"}, {Name: "CreatedAt", PKType: "", Column: "created_at"}}, PKFieldIndex: 0},
	z: new(Invoice).Values(),
}

// String returns a string representation of this struct or record.
func (s Invoice) String() string {
	res := make([]string, 7)
	res[0] = "InvoiceID: " + reform.Inspect(s.InvoiceID, true)
	res[1] = "OrderID: " + reform.Inspect(s.OrderID, true)
	res[2] = "DestinationID: " + reform.Inspect(s.DestinationID, true)
	res[3] = "SourceID: " + reform.Inspect(s.SourceID, true)
	res[4] = "Paid: " + reform.Inspect(s.Paid, true)
	res[5] = "Amount: " + reform.Inspect(s.Amount, true)
	res[6] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Invoice) Values() []interface{} {
	return []interface{}{
		s.InvoiceID,
		s.OrderID,
		s.DestinationID,
		s.SourceID,
		s.Paid,
		s.Amount,
		s.CreatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Invoice) Pointers() []interface{} {
	return []interface{}{
		&s.InvoiceID,
		&s.OrderID,
		&s.DestinationID,
		&s.SourceID,
		&s.Paid,
		&s.Amount,
		&s.CreatedAt,
	}
}

// View returns View object for that struct.
func (s *Invoice) View() reform.View {
	return InvoiceTable
}

// Table returns Table object for that record.
func (s *Invoice) Table() reform.Table {
	return InvoiceTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Invoice) PKValue() interface{} {
	return s.InvoiceID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Invoice) PKPointer() interface{} {
	return &s.InvoiceID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Invoice) HasPK() bool {
	return s.InvoiceID != InvoiceTable.z[InvoiceTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Invoice) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.InvoiceID = int64(i64)
	} else {
		s.InvoiceID = pk.(int64)
	}
}

// check interfaces
var (
	_ reform.View   = InvoiceTable
	_ reform.Struct = new(Invoice)
	_ reform.Table  = InvoiceTable
	_ reform.Record = new(Invoice)
	_ fmt.Stringer  = new(Invoice)
)

func init() {
	parse.AssertUpToDate(&InvoiceTable.s, new(Invoice))
}
