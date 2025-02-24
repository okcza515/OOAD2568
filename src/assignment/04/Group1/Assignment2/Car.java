public class Car {
    private int numSeats;
    private float engineVolume;
    private int engineMileage;
    private String transmissionType;

    public void setSeats(int numSeats) {
        this.numSeats = numSeats;
    }

    public void setEngine(float engineVolume, int engineMileage) {
        this.engineVolume = engineVolume;
        this.engineMileage = engineMileage;
    }

    public void setTransmission(String transmissionType) {
        this.transmissionType = transmissionType;
    }

    @Override
    public String toString() {
        return "Car with " + numSeats + " seats, engine volume " + engineVolume + ", engine mileage " + engineMileage + ", transmission type " + transmissionType;
    }
}
