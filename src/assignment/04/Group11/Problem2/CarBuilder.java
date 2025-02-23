public interface CarBuilder {
    void setCarType(Type carType);
    void setSeats(int seats);
    void setEngine(Engine engine);
    void setTransmission(Transmission transmission);
    void setTripComputer(TripComputer tripComputer);
    void setGPSNavigator(GPSNavigator gpsNavigator);
    Car getCar();
}
// 65070503412 Chitsanupong