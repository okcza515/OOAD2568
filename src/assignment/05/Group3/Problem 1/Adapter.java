
public class Adapter extends RoundPeg {

    private SquarePeg squarepeg;

    public Adapter(SquarePeg squarepeg) {
        this.squarepeg = squarepeg;
    }

    @Override
    public double getRadius() {
        return Math.sqrt(2) * (squarepeg.getWidth() / 2);
    }
}
