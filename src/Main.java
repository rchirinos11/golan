import wol.Wol;

class Main {
  public static void main(String[] args) {
    // Basic implementation via main method
    var mac = System.getenv("WOLMAC");
    Wol.makeMagic(mac);
    Wol.sendMagic();
  }
}
