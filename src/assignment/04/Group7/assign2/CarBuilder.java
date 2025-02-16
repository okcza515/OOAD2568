public abstract class CarBuilder {
    Car car;
  
    public void setCar(Car car) {
          this.car = car;
      }
      
      public Car getCar() {
          return car;
      }
      
      public abstract void createType();
    public abstract void createFuel();
    public abstract void createSeats();
    public abstract void createEngine();
      public abstract void createTransmission();
    public abstract void createTripComputer();
    public abstract void createGpsNavigator();
      
      public final Car buildCar() {
          Car c = new Car();
          setCar(car);
          createType();
          createFuel();
          createSeats();
          createEngine();
      createTransmission();
          createTripComputer();
      createGpsNavigator();
  
          return c;
      }
  }
  
  //Jaatupoj 65070501070