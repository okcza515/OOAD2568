// 65070501079 Pitchayuth Jampong
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
return "name";
}

public Date getDate() {
return date;
}

public String productBreakDown() {
String str = "";
for (Product product : products) {
str += product.getProductId() + " " + product.getProductName() + "\n";
}

return str;
}
}