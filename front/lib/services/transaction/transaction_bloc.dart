import 'package:bloc/bloc.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../../models/transaction_model.dart';
import 'transaction_event.dart';
import 'transaction_state.dart';
import '../../utils/config_io.dart';

// Bloc pour gérer les transactions
class TransactionBloc extends Bloc<TransactionEvent, TransactionState> {
  TransactionBloc() : super(TransactionInitial()) {
    // Associer chaque événement à une fonction de traitement
    on<CreateTransactionEvent>(_onCreateTransaction);
    on<FetchTransactionEvent>(_onFetchTransaction);
    on<FetchTransactionsByParentIdEvent>(_onFetchTransactionsByParentId);
  }

  // Fonction pour gérer l'événement de création de transaction
  void _onCreateTransaction(
      CreateTransactionEvent event, Emitter<TransactionState> emit) async {
    emit(TransactionLoading());
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/transactions'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'parent_id': event.parentId,
          'price': event.price,
          'tokens_amount': event.tokensAmount,
        }),
      );

      if (response.statusCode == 201) {
        final transactionData = jsonDecode(response.body);
        final transaction = Transaction.fromJson(transactionData);
        emit(TransactionSuccess(transaction));
      } else {
        emit(TransactionError('Erreur lors de la création de la transaction'));
      }
    } catch (e) {
      emit(TransactionError(e.toString()));
    }
  }

  // Fonction pour récupérer une transaction par ID
  void _onFetchTransaction(
      FetchTransactionEvent event, Emitter<TransactionState> emit) async {
    emit(TransactionLoading());
    try {
      final response = await http.get(
        Uri.parse('${Config.baseUrl}/transactions/${event.transactionId}'),
      );

      if (response.statusCode == 200) {
        final transactionData = jsonDecode(response.body);
        final transaction = Transaction.fromJson(transactionData);
        emit(TransactionSuccess(transaction));
      } else {
        emit(TransactionError('Transaction non trouvée'));
      }
    } catch (e) {
      emit(TransactionError(e.toString()));
    }
  }

  // Fonction pour récupérer toutes les transactions d'un parent
  void _onFetchTransactionsByParentId(FetchTransactionsByParentIdEvent event,
      Emitter<TransactionState> emit) async {
    emit(TransactionLoading());
    try {
      final response = await http.get(
        Uri.parse('${Config.baseUrl}/transactions?parent_id=${event.parentId}'),
      );

      if (response.statusCode == 200) {
        final transactionsData = jsonDecode(response.body) as List;
        final transactions = transactionsData.map((t) => Transaction.fromJson(t)).toList();
        emit(TransactionsSuccess(transactions));
      } else {
        emit(TransactionError('Aucune transaction trouvée pour ce parent'));
      }
    } catch (e) {
      emit(TransactionError(e.toString()));
    }
  }
}
