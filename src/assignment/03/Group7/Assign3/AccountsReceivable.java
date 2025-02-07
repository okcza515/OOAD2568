interface PaymentHandler {
    void chargeCustomer();
}

class PaymentService implements PaymentHandler {
    @Override
    public void chargeCustomer() {
        System.out.println("Customer has been charged.");
    }
}

public class AccountsReceivable {
	private PaymentHandler paymentHandler;
	
	public AccountsReceivable(PaymentHandler paymentHandler){
		this.paymentHandler = paymentHandler;
	}
	
	public void postPayment(){
		paymentHandler.chargeCustomer();
	}
}

//65070501053