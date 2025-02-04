import java.util.ArrayList;
import java.util.List;
public class Main {
    public static void main(String[] args) {
        
        Customer customer = new Customer("KEN");

        List<Product> products = new ArrayList<>();
        products.add(new Product(101, "Laptop"));
        products.add(new Product(102, "Mouse"));
        
        ReportTransaction Report = new ReportTransaction(customer,products);
        AccountProcess Account = new AccountProcess();
        
        System.out.println("Customer Name: " + Report.getName());
        System.out.println("Transaction Date: " + Report.getDate());
        System.out.println(Report.productBreakDown());

        Account.prepareInvoice();
        Account.chargeCustomer();
      
    }
}