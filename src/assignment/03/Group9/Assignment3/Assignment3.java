// 2. Product related interfaces and classes
//Bew
interface Product {
    String getId();
    String getName();
    String getBreakdown();
    BigDecimal getPrice();
}

class ProductImpl implements Product {
    private String id;
    private String name;
    private String breakdown;
    private BigDecimal price;

    public ProductImpl(String id, String name, String breakdown, BigDecimal price) {
        this.id = id;
        this.name = name;
        this.breakdown = breakdown;
        this.price = price;
    }

    @Override
    public String getId() { return id; }

    @Override
    public String getName() { return name; }

    @Override
    public String getBreakdown() { return breakdown; }

    @Override
    public BigDecimal getPrice() { return price; }
}

// 4. Transaction Generator interface and implementation
// Sawitt Ngamvilaisiriwong 65070503469
interface TransactionGenerator {
    Transaction generateTransaction(Customer customer, Product product) throws TransactionException;
}

class TransactionGeneratorImpl implements TransactionGenerator {
    private static final String TRANSACTION_ID_PREFIX = "TRX";
    private final AtomicLong transactionCounter = new AtomicLong(0);

    @Override
    public Transaction generateTransaction(Customer customer, Product product) throws TransactionException {
        String transactionId = TRANSACTION_ID_PREFIX + transactionCounter.incrementAndGet();
        return new CustomerTransaction(transactionId, customer, product);
    }
}

// 5. Accounts Receivable interface and implementation
interface AccountsReceivable {
    void recordTransaction(Transaction transaction) throws AccountingException;
    BigDecimal getCustomerBalance(Customer customer);
    List<Transaction> getCustomerTransactions(Customer customer);
}

class AccountsReceivableImpl implements AccountsReceivable {
    private Map<String, List<Transaction>> customerTransactions = new ConcurrentHashMap<>();
    private Map<String, BigDecimal> customerBalances = new ConcurrentHashMap<>();

    @Override
    public void recordTransaction(Transaction transaction) throws AccountingException {
        String customerId = transaction.getCustomer().getId();
        customerTransactions.computeIfAbsent(customerId, k -> new ArrayList<>())
                          .add(transaction);
        
        customerBalances.compute(customerId, (k, v) -> 
            (v == null ? BigDecimal.ZERO : v).add(transaction.getAmount()));
    }

    @Override
    public BigDecimal getCustomerBalance(Customer customer) {
        return customerBalances.getOrDefault(customer.getId(), BigDecimal.ZERO);
    }

    @Override
    public List<Transaction> getCustomerTransactions(Customer customer) {
        return customerTransactions.getOrDefault(customer.getId(), Collections.emptyList());
    }
}

// 6. Service layer to orchestrate the operations
//Wat
class TransactionService {
    private final TransactionGenerator transactionGenerator;
    private final AccountsReceivable accountsReceivable;

    public TransactionService(TransactionGenerator transactionGenerator, 
                            AccountsReceivable accountsReceivable) {
        this.transactionGenerator = transactionGenerator;
        this.accountsReceivable = accountsReceivable;
    }

    public void processTransaction(Customer customer, Product product) 
            throws TransactionException, AccountingException {
        Transaction transaction = transactionGenerator.generateTransaction(customer, product);
        transaction.execute();
        accountsReceivable.recordTransaction(transaction);
    }

    public BigDecimal getCustomerBalance(Customer customer) {
        return accountsReceivable.getCustomerBalance(customer);
    }
}

// 7. Custom exceptions
class TransactionException extends Exception {
    public TransactionException(String message) {
        super(message);
    }
}

class AccountingException extends Exception {
    public AccountingException(String message) {
        super(message);
    }
}