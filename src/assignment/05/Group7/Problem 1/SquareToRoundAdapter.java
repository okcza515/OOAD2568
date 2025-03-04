package assignment5.problem1;

public class SquareToRoundAdapter implements PegAdapter {
    public static RoundPeg SquareToRound(SquarePeg peg) {
        return new RoundPeg((Math.sqrt(Math.pow((peg.getWidth() / 2), 2) * 2)));
    }
}
