abstract class ParentEvent {}

class CreateParent extends ParentEvent {
  final int userId;
  final int tokensAmount;

  CreateParent({required this.userId, required this.tokensAmount});
}

class GetParentById extends ParentEvent {
  final int id;

  GetParentById(this.id);
}

class UpdateParent extends ParentEvent {
  final int id;
  final int tokensAmount;

  UpdateParent({required this.id, required this.tokensAmount});
}

class DeleteParent extends ParentEvent {
  final int id;

  DeleteParent(this.id);
}
