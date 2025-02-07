import java.util.Date;
import java.util.List;
public class Main {
    public static void main(String[] args) {
        // Create a customer
        Customer customer = new Customer();
        customer.setName("Diw");

        // Create products
        List<Product> products = new ArrayList<>();
        Product product1 = new Product();
        product1.setProductId(101);
        product1.setProductName("Laptop");
        
        Product product2 = new Product();
        product2.setProductId(102);
        product2.setProductName("Mouse");
        
        products.add(product1);
        products.add(product2);
        
        // Create CustomerTransaction
        CustomerTransaction transaction = new CustomerTransaction(customer, products);
        
        // Create AccountsReceivable
        AccountsReceivable account = new AccountsReceivable(transaction);
        
        // Print transaction details
        System.out.println("Customer Name: " + transaction.getName());
        System.out.println("Transaction Date: " + transaction.getDate());
        System.out.println(transaction.productBreakDown());

        // Process the transaction
        account.sendInvoice();
        account.postPayment();
    }
}

//Intouch Krajangprateep 65070503442