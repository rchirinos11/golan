package wol;

import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Wol {
  private static final byte FF = (byte) 0xFF;
  private static final byte[] magic = new byte[102];

  public static void makeMagic(final String mac) {
    System.out.println("Creating magic packet...");

    final String[] splits = mac.split(":");
    for (int i = 0; i < magic.length; i++) {
      magic[i] = (i < 6) ? FF : (byte) Integer.parseInt(splits[i % 6], 16);
    }
  }

  public static void sendMagic() {
    System.out.println("Sending magic packet...");

    try {
      final var socket = new DatagramSocket();
      final var address = InetAddress.getByName("255.255.255.255");
      final var packet = new DatagramPacket(magic, magic.length, address, 9);
      socket.send(packet);
      System.out.println("Magic packet sent.");
      socket.close();
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  static void printMagic() {
    final String hexString = IntStream.range(0, magic.length)
        .mapToObj(i -> String.format("%02X", magic[i]))
        .collect(Collectors.joining(":"));
    System.out.println(hexString);
  }
}
