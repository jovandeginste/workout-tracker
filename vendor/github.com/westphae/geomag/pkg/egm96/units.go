package egm96

// Factors for converting between radians and degrees
// and between meters and feet.
const (
	Deg = 1/57.29577951308232 // number of radians per degree
	Ft  = 0.3048              // number of meters per foot
)

// DMSToDegrees converts integral degrees d, minutes m and seconds s (all of type float64)
// to a float-valued degrees amount.
//
// If d<0 then must pass m>0 and s>0;
// if d==0 and m<0 then must pass s>0.
func DMSToDegrees(d, m, s float64) (dd float64) {
	var sgn float64 = 1
	if d<0 {
		sgn = -1
		d = -d
	}
	if d==0 && m<0 {
		sgn = -1
		m = -m
	}
	return sgn*(d+(m+s/60)/60)
}

// DegreesToDMS converts float degrees dd to
// whole degrees d, minutes m and decimal seconds s (all of type float64).
func DegreesToDMS(dd float64) (d, m, s float64) {
	var sgn float64 = 1
	if dd<0 {
		sgn = -1
		dd = -dd
	}
	d = float64(int(dd))
	z := (dd-d)*60
	m = float64(int(z))
	s = (z-m)*60
	if sgn==-1 && d==0 && m==0 {
		return 0, 0, -s
	}
	if sgn==-1 && d==0 {
		return 0, -m, s
	}
	return sgn*d, m, s
}
