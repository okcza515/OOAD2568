
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

    public String getName() {
        return this.customer.getName();
    }

    public Date getDate() {
        return this.date;
    }

    public String productBreakDown() {
      String str = "";
      for (int i = 0; i < this.products.size(); i++) {
          Product product = this.products.get(i);
          str += String.format("* ID:%d - %s\n", product.getProductId(), product.getProductName());
      }
      return str;
    }
}
