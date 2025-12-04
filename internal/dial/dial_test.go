package dial_test


import(
	"testing"
	"strconv"
	"veralfre.com/aoc-2025/internal/dial"
)


func TestNewDial(t *testing.T) {
	d:= dial.NewDial(100)
	if d == nil{
		t.Error("Expected dial to be created, got nil")
	}

	if d.GetCurrentPosition() != 50{
		t.Errorf("Expected starting position to be 50, got %d", d.GetCurrentPosition())
	}
}

func TestDial_Turn(t *testing.T) {
	d:= dial.NewDial(100)

	d.Turn("R", 30)
	if d.GetCurrentPosition() != 80{
		t.Errorf("Expected position to be 80 after turning right 30 from 50, got %d", d.GetCurrentPosition())
	}

	d.Turn("L", 50)
	if d.GetCurrentPosition() != 30{
		t.Errorf("Expected position to be 30 after turning left 50 from 80, got %d", d.GetCurrentPosition())
	}

	d.Turn("R", 80)
	if d.GetCurrentPosition() != 10{
		t.Errorf("Expected position to be 10 after turning right 80 from 30, got %d", d.GetCurrentPosition())
	}

	d.Turn("L", 20)
	if d.GetCurrentPosition() != 90{
		t.Errorf("Expected position to be 90 after turning left 20 from 10, got %d", d.GetCurrentPosition())
	}
}

func TestDialComplex(t *testing.T){
	d:= dial.NewDial(100)
	directions := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	for _, direction := range directions{
		dir := string(direction[0])
		steps, _ := strconv.Atoi(direction[1:])
		d.Turn(dir, steps)
	}
	if d.GetCurrentPosition() != 32 {
		t.Errorf("Expected final position to be 82, got %d", d.GetCurrentPosition())
	}
}

func TestSimplePassword(t *testing.T){
	d:= dial.NewDial(100)
	directions := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	password := d.GetSimplePassword(directions)
	if password != 3 {
		t.Errorf("Expected simple password to be 3, got %d", password)
	}
}

func TestDialTurnOver(t *testing.T){
	d:= dial.NewDial(100)

	turns := d.Turn("R", 250)
	if turns != 2{
		t.Errorf("[%d]Expected 2 turn overs after turning right 250 on a 100 step dial, got %d", d.GetCurrentPosition(),turns)
	}

	turns = d.Turn("L", 450)
	if turns != 4{
		t.Errorf("[%d]Expected 4 turn overs after turning left 450 on a 100 step dial, got %d", d.GetCurrentPosition(), turns)
	}
}

func TestComplexPassword(t *testing.T){
	d:= dial.NewDial(100)
	directions := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	password := d.GetComplexPassword(directions)
	if password != 6 {
		t.Errorf("Expected complex password to be 6, got %d", password)
	}

	d.Reset()
	directions = []string{"R1000", "L1000", "L50"}
	password = d.GetComplexPassword(directions)
	if password != 21 {
		t.Errorf("Expected complex password to be 21, got %d", password)
	}
	// R1000 # +10 (50) 10
	// L1000 # +10 (50) 20
	// L50   # +1  (0)  21
	// R1    # +0  (1)  21
	// L1    # +1  (0)  22
	// L1    # +0  (99) 22
	// R1    # +1  (0)  23
	// R100  # +1  (0)  24
	// R1    # +0  (1)  24
	d.Reset()
	directions = []string{"R1000", "L1000", "L50", "R1", "L1", "L1", "R1", "R100", "R1"}
	password = d.GetComplexPassword(directions)
	if password != 24 {
		t.Errorf("Expected complex password to be 24, got %d", password)
	}
}

func TestDialTurnHighSteps(t *testing.T){
	d:= dial.NewDial(100)

	turns := d.Turn("R", 1000)
	if turns != 10 {
		t.Errorf("Expected 10 turn overs after turning right 1000 on a 100 step dial, got %d", turns)
	}

	turns += d.Turn("L", 1000)
	if turns != 20 {
		t.Errorf("Expected 20 turn overs after turning left 1000 on a 100 step dial, got %d", turns)
	}

}