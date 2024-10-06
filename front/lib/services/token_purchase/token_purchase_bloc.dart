import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/services/token_purchase/token_purchase_event.dart';
import 'package:front/services/token_purchase/token_purchase_state.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import '../../utils/config_io.dart';

// Bloc
class TokenPurchaseBloc extends Bloc<TokenPurchaseEvent, TokenPurchaseState> {
  TokenPurchaseBloc() : super(TokenPurchaseInitial());

  Stream<TokenPurchaseState> mapEventToState(TokenPurchaseEvent event) async* {
    if (event is PurchaseTokens) {
      yield TokenPurchaseLoading();
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

        final paymentIntentData = jsonDecode(response.body);

        yield TokenPurchaseSuccess(paymentIntentData['id']);
      } catch (e) {
        yield TokenPurchaseError(e.toString());
      }
    }
  }
}
