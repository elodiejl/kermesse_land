class Student {
  final int id;
  final String name;
  final int parentId;
  final int tokenAmount;

  Student( {required this.id, required this.name, required this.parentId, required this.tokenAmount});

  factory Student.fromJson(Map<String, dynamic> json) {
    return Student(
      id: json['id'],
      name: json['user']['username'],
      parentId: json['parent_id'],
      tokenAmount: json['token_amount']
    );
  }
}
