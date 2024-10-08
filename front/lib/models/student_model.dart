class Student {
  final int id;
  final String name;
  final int parentId;

  Student({required this.id, required this.name, required this.parentId});

  factory Student.fromJson(Map<String, dynamic> json) {
    return Student(
      id: json['id'],
      name: json['username'],
      parentId: json['parent_id'],
    );
  }
}
