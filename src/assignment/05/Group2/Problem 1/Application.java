
public class Application {

	public static void main(String[] args) {
		RoundHole hole = new RoundHole(5);
        RoundPeg rpeg = new RoundPeg(5);
        if (hole.fits(rpeg)) {
            System.out.println("Round peg r5 fits round hole r5.");
        }
        
// The following code is used after add SquarePeg
        
        SquarePeg smallSqPeg = new SquarePeg(2);
        SquarePeg largeSqPeg = new SquarePeg(20);
        
        Adapter smallAdater = new Adapter(smallSqPeg);
        Adapter largeAdapter = new Adapter(largeSqPeg);

        if (hole.fits(smallAdater)) {
            System.out.println("Square peg w2 fits round hole r5.");
        }
        if (!hole.fits(largeAdapter)) {
            System.out.println("Square peg w20 does not fit into round hole r5.");
        }
	}
}

//Korawit Sritotum 65070503402