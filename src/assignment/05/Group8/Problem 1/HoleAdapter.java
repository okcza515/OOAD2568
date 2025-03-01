public class HoleAdapter extends RoundPeg {
    private SquarePeg squarePeg;

    public HoleAdapter(SquarePeg squarePeg){
        this.squarePeg = squarePeg;
    }
    @Override
    public double getRadius(){
        return (Math.sqrt(Math.pow((squarePeg.getWidth() / 2), 2) * 2));
    }
}
