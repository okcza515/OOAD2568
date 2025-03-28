
public class SportCarBuilder extends CarBuilder {

    public SportCarBuilder() {
        this.carType = Type.SPORTS_CAR;
        this.tripComputer = new TripComputer();
        this.gpsNavigator = new GPSNavigator();
    }

}

/*
 _   __                                         _____           _                 _ _                       _           _   
| | / /                                        /  ___|         | |               | | |                     | |         | |  
| |/ /  __ _ _ __   __ _ ___  ___  _ __ _ __   \ `--. _   _  __| |_   _  ___   __| | |__  _   _ _ __  _ __ | |__   ___ | |_ 
|    \ / _` | '_ \ / _` / __|/ _ \| '__| '_ \   `--. \ | | |/ _` | | | |/ _ \ / _` | '_ \| | | | '_ \| '_ \| '_ \ / _ \| __|
| |\  \ (_| | | | | (_| \__ \ (_) | |  | | | | /\__/ / |_| | (_| | |_| | (_) | (_| | |_) | |_| | | | | |_) | | | | (_) | |_ 
\_| \_/\__,_|_| |_|\__,_|___/\___/|_|  |_| |_| \____/ \__,_|\__,_|\__, |\___/ \__,_|_.__/ \__,_|_| |_| .__/|_| |_|\___/ \__|
                                                                   __/ |                             | |                    
                                                                  |___/                              |_|                    
  ____  _____  _____  ___________ _____  _____  __  _____  ____  ______                                                     
 / ___||  ___||  _  ||___  /  _  |  ___||  _  |/  ||  _  |/ ___||___  /                                                     
/ /___ |___ \ | |/' |   / /| |/' |___ \ | |/' |`| || |/' / /___    / /                                                      
| ___ \    \ \|  /| |  / / |  /| |   \ \|  /| | | ||  /| | ___ \  / /                                                       
| \_/ |/\__/ /\ |_/ /./ /  \ |_/ /\__/ /\ |_/ /_| |\ |_/ / \_/ |./ /                                                        
\_____/\____/  \___/ \_/    \___/\____/  \___/ \___/\___/\_____/\_/  
                                                       
 */