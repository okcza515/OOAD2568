public class AccountsReceivable {
    private PaymentService paymentService;
    private InvoiceService invoiceService;
    
    public AccountsReceivable(PaymentService paymentService, InvoiceService invoiceService) {
        this.paymentService = paymentService;
        this.invoiceService = invoiceService;
    }
    
    public void postPayment(CustomerTransaction transaction) {
        paymentService.charge(transaction);
    }
    
    public void sendInvoice(CustomerTransaction transaction) {
        Invoice invoice = invoiceService.prepareInvoice(transaction);
        System.out.println("Invoice sent: " + invoice);
    }
}
