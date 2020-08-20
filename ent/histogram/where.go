// Code generated by entc, DO NOT EDIT.

package histogram

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/gobench-io/gobench/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTime), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// Min applies equality check predicate on the "min" field. It's identical to MinEQ.
func Min(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMin), v))
	})
}

// Max applies equality check predicate on the "max" field. It's identical to MaxEQ.
func Max(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMax), v))
	})
}

// Mean applies equality check predicate on the "mean" field. It's identical to MeanEQ.
func Mean(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMean), v))
	})
}

// Stddev applies equality check predicate on the "stddev" field. It's identical to StddevEQ.
func Stddev(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStddev), v))
	})
}

// Median applies equality check predicate on the "median" field. It's identical to MedianEQ.
func Median(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedian), v))
	})
}

// P75 applies equality check predicate on the "p75" field. It's identical to P75EQ.
func P75(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP75), v))
	})
}

// P95 applies equality check predicate on the "p95" field. It's identical to P95EQ.
func P95(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP95), v))
	})
}

// P99 applies equality check predicate on the "p99" field. It's identical to P99EQ.
func P99(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP99), v))
	})
}

// P999 applies equality check predicate on the "p999" field. It's identical to P999EQ.
func P999(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP999), v))
	})
}

// WID applies equality check predicate on the "wID" field. It's identical to WIDEQ.
func WID(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWID), v))
	})
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTime), v))
	})
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTime), v))
	})
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTime), v...))
	})
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTime), v...))
	})
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTime), v))
	})
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTime), v))
	})
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTime), v))
	})
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTime), v))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// MinEQ applies the EQ predicate on the "min" field.
func MinEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMin), v))
	})
}

// MinNEQ applies the NEQ predicate on the "min" field.
func MinNEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMin), v))
	})
}

// MinIn applies the In predicate on the "min" field.
func MinIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMin), v...))
	})
}

// MinNotIn applies the NotIn predicate on the "min" field.
func MinNotIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMin), v...))
	})
}

// MinGT applies the GT predicate on the "min" field.
func MinGT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMin), v))
	})
}

// MinGTE applies the GTE predicate on the "min" field.
func MinGTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMin), v))
	})
}

// MinLT applies the LT predicate on the "min" field.
func MinLT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMin), v))
	})
}

// MinLTE applies the LTE predicate on the "min" field.
func MinLTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMin), v))
	})
}

// MaxEQ applies the EQ predicate on the "max" field.
func MaxEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMax), v))
	})
}

// MaxNEQ applies the NEQ predicate on the "max" field.
func MaxNEQ(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMax), v))
	})
}

// MaxIn applies the In predicate on the "max" field.
func MaxIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMax), v...))
	})
}

// MaxNotIn applies the NotIn predicate on the "max" field.
func MaxNotIn(vs ...int64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMax), v...))
	})
}

// MaxGT applies the GT predicate on the "max" field.
func MaxGT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMax), v))
	})
}

// MaxGTE applies the GTE predicate on the "max" field.
func MaxGTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMax), v))
	})
}

// MaxLT applies the LT predicate on the "max" field.
func MaxLT(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMax), v))
	})
}

// MaxLTE applies the LTE predicate on the "max" field.
func MaxLTE(v int64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMax), v))
	})
}

// MeanEQ applies the EQ predicate on the "mean" field.
func MeanEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMean), v))
	})
}

// MeanNEQ applies the NEQ predicate on the "mean" field.
func MeanNEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMean), v))
	})
}

// MeanIn applies the In predicate on the "mean" field.
func MeanIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMean), v...))
	})
}

// MeanNotIn applies the NotIn predicate on the "mean" field.
func MeanNotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMean), v...))
	})
}

// MeanGT applies the GT predicate on the "mean" field.
func MeanGT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMean), v))
	})
}

// MeanGTE applies the GTE predicate on the "mean" field.
func MeanGTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMean), v))
	})
}

// MeanLT applies the LT predicate on the "mean" field.
func MeanLT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMean), v))
	})
}

// MeanLTE applies the LTE predicate on the "mean" field.
func MeanLTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMean), v))
	})
}

// StddevEQ applies the EQ predicate on the "stddev" field.
func StddevEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStddev), v))
	})
}

// StddevNEQ applies the NEQ predicate on the "stddev" field.
func StddevNEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStddev), v))
	})
}

// StddevIn applies the In predicate on the "stddev" field.
func StddevIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStddev), v...))
	})
}

// StddevNotIn applies the NotIn predicate on the "stddev" field.
func StddevNotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStddev), v...))
	})
}

// StddevGT applies the GT predicate on the "stddev" field.
func StddevGT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStddev), v))
	})
}

// StddevGTE applies the GTE predicate on the "stddev" field.
func StddevGTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStddev), v))
	})
}

// StddevLT applies the LT predicate on the "stddev" field.
func StddevLT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStddev), v))
	})
}

// StddevLTE applies the LTE predicate on the "stddev" field.
func StddevLTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStddev), v))
	})
}

// MedianEQ applies the EQ predicate on the "median" field.
func MedianEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedian), v))
	})
}

// MedianNEQ applies the NEQ predicate on the "median" field.
func MedianNEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMedian), v))
	})
}

// MedianIn applies the In predicate on the "median" field.
func MedianIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMedian), v...))
	})
}

// MedianNotIn applies the NotIn predicate on the "median" field.
func MedianNotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMedian), v...))
	})
}

// MedianGT applies the GT predicate on the "median" field.
func MedianGT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMedian), v))
	})
}

// MedianGTE applies the GTE predicate on the "median" field.
func MedianGTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMedian), v))
	})
}

// MedianLT applies the LT predicate on the "median" field.
func MedianLT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMedian), v))
	})
}

// MedianLTE applies the LTE predicate on the "median" field.
func MedianLTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMedian), v))
	})
}

// P75EQ applies the EQ predicate on the "p75" field.
func P75EQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP75), v))
	})
}

// P75NEQ applies the NEQ predicate on the "p75" field.
func P75NEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldP75), v))
	})
}

// P75In applies the In predicate on the "p75" field.
func P75In(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldP75), v...))
	})
}

// P75NotIn applies the NotIn predicate on the "p75" field.
func P75NotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldP75), v...))
	})
}

// P75GT applies the GT predicate on the "p75" field.
func P75GT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldP75), v))
	})
}

// P75GTE applies the GTE predicate on the "p75" field.
func P75GTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldP75), v))
	})
}

// P75LT applies the LT predicate on the "p75" field.
func P75LT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldP75), v))
	})
}

// P75LTE applies the LTE predicate on the "p75" field.
func P75LTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldP75), v))
	})
}

// P95EQ applies the EQ predicate on the "p95" field.
func P95EQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP95), v))
	})
}

// P95NEQ applies the NEQ predicate on the "p95" field.
func P95NEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldP95), v))
	})
}

// P95In applies the In predicate on the "p95" field.
func P95In(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldP95), v...))
	})
}

// P95NotIn applies the NotIn predicate on the "p95" field.
func P95NotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldP95), v...))
	})
}

// P95GT applies the GT predicate on the "p95" field.
func P95GT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldP95), v))
	})
}

// P95GTE applies the GTE predicate on the "p95" field.
func P95GTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldP95), v))
	})
}

// P95LT applies the LT predicate on the "p95" field.
func P95LT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldP95), v))
	})
}

// P95LTE applies the LTE predicate on the "p95" field.
func P95LTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldP95), v))
	})
}

// P99EQ applies the EQ predicate on the "p99" field.
func P99EQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP99), v))
	})
}

// P99NEQ applies the NEQ predicate on the "p99" field.
func P99NEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldP99), v))
	})
}

// P99In applies the In predicate on the "p99" field.
func P99In(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldP99), v...))
	})
}

// P99NotIn applies the NotIn predicate on the "p99" field.
func P99NotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldP99), v...))
	})
}

// P99GT applies the GT predicate on the "p99" field.
func P99GT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldP99), v))
	})
}

// P99GTE applies the GTE predicate on the "p99" field.
func P99GTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldP99), v))
	})
}

// P99LT applies the LT predicate on the "p99" field.
func P99LT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldP99), v))
	})
}

// P99LTE applies the LTE predicate on the "p99" field.
func P99LTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldP99), v))
	})
}

// P999EQ applies the EQ predicate on the "p999" field.
func P999EQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldP999), v))
	})
}

// P999NEQ applies the NEQ predicate on the "p999" field.
func P999NEQ(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldP999), v))
	})
}

// P999In applies the In predicate on the "p999" field.
func P999In(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldP999), v...))
	})
}

// P999NotIn applies the NotIn predicate on the "p999" field.
func P999NotIn(vs ...float64) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldP999), v...))
	})
}

// P999GT applies the GT predicate on the "p999" field.
func P999GT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldP999), v))
	})
}

// P999GTE applies the GTE predicate on the "p999" field.
func P999GTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldP999), v))
	})
}

// P999LT applies the LT predicate on the "p999" field.
func P999LT(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldP999), v))
	})
}

// P999LTE applies the LTE predicate on the "p999" field.
func P999LTE(v float64) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldP999), v))
	})
}

// WIDEQ applies the EQ predicate on the "wID" field.
func WIDEQ(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWID), v))
	})
}

// WIDNEQ applies the NEQ predicate on the "wID" field.
func WIDNEQ(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWID), v))
	})
}

// WIDIn applies the In predicate on the "wID" field.
func WIDIn(vs ...string) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWID), v...))
	})
}

// WIDNotIn applies the NotIn predicate on the "wID" field.
func WIDNotIn(vs ...string) predicate.Histogram {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Histogram(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWID), v...))
	})
}

// WIDGT applies the GT predicate on the "wID" field.
func WIDGT(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWID), v))
	})
}

// WIDGTE applies the GTE predicate on the "wID" field.
func WIDGTE(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWID), v))
	})
}

// WIDLT applies the LT predicate on the "wID" field.
func WIDLT(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWID), v))
	})
}

// WIDLTE applies the LTE predicate on the "wID" field.
func WIDLTE(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWID), v))
	})
}

// WIDContains applies the Contains predicate on the "wID" field.
func WIDContains(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWID), v))
	})
}

// WIDHasPrefix applies the HasPrefix predicate on the "wID" field.
func WIDHasPrefix(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWID), v))
	})
}

// WIDHasSuffix applies the HasSuffix predicate on the "wID" field.
func WIDHasSuffix(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWID), v))
	})
}

// WIDEqualFold applies the EqualFold predicate on the "wID" field.
func WIDEqualFold(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWID), v))
	})
}

// WIDContainsFold applies the ContainsFold predicate on the "wID" field.
func WIDContainsFold(v string) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWID), v))
	})
}

// HasMetric applies the HasEdge predicate on the "metric" edge.
func HasMetric() predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetricTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetricTable, MetricColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetricWith applies the HasEdge predicate on the "metric" edge with a given conditions (other predicates).
func HasMetricWith(preds ...predicate.Metric) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetricInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetricTable, MetricColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Histogram) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Histogram) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Histogram) predicate.Histogram {
	return predicate.Histogram(func(s *sql.Selector) {
		p(s.Not())
	})
}
