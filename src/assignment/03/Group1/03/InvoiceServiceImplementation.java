public class InvoiceServiceImplementation implements InvoiceService {
    @Override
    public Invoice prepareInvoice(CustomerTransaction transaction) {
        System.out.println("Preparing invoice for customer " + transaction.getCustomer().getName());
        return new Invoice(transaction);
    }
}
