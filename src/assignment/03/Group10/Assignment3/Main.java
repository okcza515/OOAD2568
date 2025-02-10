// 65070501076 Danai Saengbuamad
import java.util.ArrayList;

public class Main {
    public static void main(String[] args){
        Customer customer = new Customer("Chayapol");
        List<Product> products = new ArrayList<Product>();
        products.add(new Product(1, "T-Shirt"));
        products.add(new Product(2, "Jeans"));

        CustomerTransaction transaction = new CustomerTransaction(customer, products);

        System.out.println(transaction.getName());
        System.out.println(transaction.productBreakDown());
        System.out.println(transaction.getDate());

        AccountsReceivable accountsReceivable = new AccountsReceivable(transaction);

        accountsReceivable.sendInvoice();
        accountsReceivable.postPayment();

        ReportGenerator report = new ReportGenerator(transaction);
        report.generateReport();
    }
}
