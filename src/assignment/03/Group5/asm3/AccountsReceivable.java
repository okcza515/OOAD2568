class PaymentService implements PaymentHandler {
	@Override
	public void chargeCustomer() {
		System.out.println("Customer has been charged.");
	}
}

class InvoiceService implements InvoiceHandler {
	@Override
	public void prepareInvoice() {
		System.out.println("Invoice has been prepared.");
	}
}

public class AccountsReceivable {
	private PaymentHandler paymentHandler;
	private InvoiceHandler invoiceHandler;
	
	public AccountsReceivable(PaymentHandler paymentHandler, InvoiceHandler invoiceHandler){
		this.paymentHandler = paymentHandler;
		this.invoiceHandler = invoiceHandler;
	}
	
	public void postPayment(){
		paymentHandler.chargeCustomer();
	}

	public void sendInvoice(){
		invoiceHandler.prepareInvoice();
		// sends the invoice
	}
}

//65070501074 Napat Sinjindawong Group 5
