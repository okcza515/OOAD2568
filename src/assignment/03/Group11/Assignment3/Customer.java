public class Customer {
    private String name;

    // Default constructor
    public Customer() {}

    // Constructor with name
    public Customer(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    // Override toString() for better debugging
    @Override
    public String toString() {
        return "Customer{name='" + name + "'}";
    }
}

//Intouch 65070503442