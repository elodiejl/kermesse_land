class Kermesse {
  final int id;
  final String name;
  final String location;
  final String date;

  Kermesse({
    required this.id,
    required this.name,
    required this.location,
    required this.date,
  });

  factory Kermesse.fromJson(Map<String, dynamic> json) {
    return Kermesse(
      id: json['id'],
      name: json['name'],
      location: json['location'],
      date: json['date'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'location': location,
      'date': date,
    };
  }
}

class KermesseUpdate {
  final int? id;
  final String? name;
  final String? location;
  final String? date;

  KermesseUpdate({
    this.id,
    this.name,
    this.location,
    this.date,
  });

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'location': location,
      'date': date,
    };
  }
}
