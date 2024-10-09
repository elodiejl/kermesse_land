class Stand {
  final int id;
  final String name;
  final String standType;
  final int participationCost;
  final int teneurId;
  final int kermesseId;
  final int stock;

  Stand({
    required this.id,
    required this.name,
    required this.standType,
    required this.participationCost,
    required this.teneurId,
    required this.kermesseId,
    required this.stock,
  });

  factory Stand.fromJson(Map<String, dynamic> json) {
    return Stand(
      id: json['id'],
      name: json['name'],
      standType: json['stand_type'],
      participationCost: json['participation_cost'],
      teneurId: json['teneur_id'],
      kermesseId: json['kermesse_id'],
      stock: json['stock'],
    );
  }
}
