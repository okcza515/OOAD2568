public abstract class DataSourceDecorator implements DataSource {
    private DataSource wrappee;

    public DataSourceDecorator(DataSource source) {
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(data);
    }

    @Override
    public String readData() {
        return wrappee.readData();
    }
}







// /$$$$$$$$ /$$                                               /$$                 /$$                                                                  
// |__  $$__/| $$                                              | $$                | $$                                                                  
// | $$   | $$$$$$$   /$$$$$$  /$$$$$$$   /$$$$$$   /$$$$$$ | $$$$$$$   /$$$$$$ | $$                                                                  
// | $$   | $$__  $$ |____  $$| $$__  $$ |____  $$ /$$__  $$| $$__  $$ /$$__  $$| $$                                                                  
// | $$   | $$  \ $$  /$$$$$$$| $$  \ $$  /$$$$$$$| $$  \ $$| $$  \ $$| $$  \ $$| $$                                                                  
// | $$   | $$  | $$ /$$__  $$| $$  | $$ /$$__  $$| $$  | $$| $$  | $$| $$  | $$| $$                                                                  
// | $$   | $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$$| $$$$$$$/| $$  | $$|  $$$$$$/| $$                                                                  
// |__/   |__/  |__/ \_______/|__/  |__/ \_______/| $$____/ |__/  |__/ \______/ |__/                                                                  
//                                                 | $$                                                                                                
//                                                 | $$                                                                                                
//                                                 |__/                                                                                                
// /$$$$$$$$ /$$                                       /$$     /$$                                                                       /$$            
// |__  $$__/| $$                                      | $$    | $$                                                                      | $$            
// | $$   | $$$$$$$   /$$$$$$  /$$$$$$$   /$$$$$$  /$$$$$$  | $$$$$$$   /$$$$$$  /$$  /$$  /$$  /$$$$$$   /$$$$$$   /$$$$$$$ /$$   /$$| $$   /$$      
// | $$   | $$__  $$ |____  $$| $$__  $$ /$$__  $$|_  $$_/  | $$__  $$ |____  $$| $$ | $$ | $$ /$$__  $$ /$$__  $$ /$$_____/| $$  | $$| $$  /$$/      
// | $$   | $$  \ $$  /$$$$$$$| $$  \ $$| $$  \ $$  | $$    | $$  \ $$  /$$$$$$$| $$ | $$ | $$| $$$$$$$$| $$$$$$$$|  $$$$$$ | $$  | $$| $$$$$$/       
// | $$   | $$  | $$ /$$__  $$| $$  | $$| $$  | $$  | $$ /$$| $$  | $$ /$$__  $$| $$ | $$ | $$| $$_____/| $$_____/ \____  $$| $$  | $$| $$_  $$       
// | $$   | $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$$  |  $$$$/| $$  | $$|  $$$$$$$|  $$$$$/$$$$/|  $$$$$$$|  $$$$$$$ /$$$$$$$/|  $$$$$$/| $$ \  $$      
// |__/   |__/  |__/ \_______/|__/  |__/ \____  $$   \___/  |__/  |__/ \_______/ \_____/\___/  \_______/ \_______/|_______/  \______/ |__/  \__/      
//                                         /$$  \ $$                                                                                                    
//                                         |  $$$$$$/                                                                                                    
//                                         \______/                                                                                                     
// /$$$$$$  /$$$$$$$   /$$$$$$  /$$$$$$$$ /$$$$$$  /$$$$$$$   /$$$$$$    /$$    /$$$$$$   /$$$$$$   /$$$$$$                                            
// /$$__  $$| $$____/  /$$$_  $$|_____ $$//$$$_  $$| $$____/  /$$$_  $$ /$$$$   /$$$_  $$ /$$__  $$ /$$__  $$                                           
// | $$  \__/| $$      | $$$$\ $$     /$$/| $$$$\ $$| $$      | $$$$\ $$|_  $$  | $$$$\ $$|__/  \ $$|__/  \ $$                                           
// | $$$$$$$ | $$$$$$$ | $$ $$ $$    /$$/ | $$ $$ $$| $$$$$$$ | $$ $$ $$  | $$  | $$ $$ $$  /$$$$$$/   /$$$$$/                                           
// | $$__  $$|_____  $$| $$\ $$$$   /$$/  | $$\ $$$$|_____  $$| $$\ $$$$  | $$  | $$\ $$$$ /$$____/   |___  $$                                           
// | $$  \ $$ /$$  \ $$| $$ \ $$$  /$$/   | $$ \ $$$ /$$  \ $$| $$ \ $$$  | $$  | $$ \ $$$| $$       /$$  \ $$                                           
// |  $$$$$$/|  $$$$$$/|  $$$$$$/ /$$/    |  $$$$$$/|  $$$$$$/|  $$$$$$/ /$$$$$$|  $$$$$$/| $$$$$$$$|  $$$$$$/                                           
// \______/  \______/  \______/ |__/      \______/  \______/  \______/ |______/ \______/ |________/ \______/                                            