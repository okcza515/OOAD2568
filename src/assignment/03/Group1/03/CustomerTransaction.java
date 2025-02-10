import java.util.Date;
import java.util.List;

public class CustomerTransaction {
    private List<Product> products;
    private Customer customer;
    private Date date;

    public CustomerTransaction(Customer customer, List<Product> products) {
        this.products = products;
        this.customer = customer;
        this.date = new Date(); 
    }

    public List<Product> getProducts() {
        return products;
    }

    public Customer getCustomer() {
        return customer;
    }

    public Date getDate() {
        return date;
    }
}
