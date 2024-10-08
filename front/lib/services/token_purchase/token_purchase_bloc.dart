import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:kermesse_land/services/token_purchase/token_purchase_event.dart';
import 'package:kermesse_land/services/token_purchase/token_purchase_state.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import '../../utils/config_io.dart';

class TokenPurchaseBloc extends Bloc<TokenPurchaseEvent, TokenPurchaseState> {
  TokenPurchaseBloc() : super(TokenPurchaseInitial()) {
    on<PurchaseTokens>((event, emit) async {
      emit(TokenPurchaseLoading());
      try {
        final response = await http.post(
          Uri.parse('${Config.baseUrl}/api/create-payment-intent'),
          headers: {'Content-Type': 'application/json'},
          body: jsonEncode({
            'amount': event.tokenAmount * 100, // Convertir en centimes
            'currency': 'eur',
            'parent_id': event.parentId,
          }),
        );

        if (response.statusCode == 200) {
          final paymentIntentData = jsonDecode(response.body);
          emit(TokenPurchaseSuccess(paymentIntentData['id']));
        } else {
          emit(TokenPurchaseError('Erreur lors de la création du paiement.'));
        }
      } catch (e) {
        emit(TokenPurchaseError(e.toString()));
      }
    });
  }
}
