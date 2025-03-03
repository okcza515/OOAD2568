public class Adapter extends RoundPeg {
    private SquarePeg squarePeg;

    public Adapter(SquarePeg squarePeg) {
        this.squarePeg = squarePeg;
    }

    @Override
    public double getRadius() {
        return Math.sqrt(Math.pow((squarePeg.getWidth() / 2), 2) * 2);
    }
}

//Napat Sinjindawong 65070501074