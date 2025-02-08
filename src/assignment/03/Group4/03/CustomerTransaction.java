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
    return customer.getName();
  }

  public Date getDate() {
    return this.date;
  }

  public String productBreakDown() {
    StringBuilder print = new StringBuilder();
    for (Product product : products) {
      print.append(product.getProductId()).append(" ").append(product.getProductName()).append("\n");
    }
    return print.toString();
  }
}
