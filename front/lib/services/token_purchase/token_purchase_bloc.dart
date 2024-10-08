import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_stripe/flutter_stripe.dart';
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
          headers: {'Authorization': 'Bearer ${event.token}', 'Content-Type': 'application/json'},
          body: jsonEncode({
            'amount': event.tokenAmount * 100, // Convertir en centimes
            'currency': 'eur',
            'parent_id': event.parentId,
          }),
        );

        if (response.statusCode == 200) {
          final paymentIntentData = jsonDecode(response.body);
          final clientSecret = paymentIntentData['clientSecret'];

          // Appelez la méthode pour afficher le formulaire de paiement
          await showPaymentSheet(clientSecret, emit, event.parentId, event.tokenAmount, event.token);
          //emit(TokenPurchaseSuccess(paymentIntentData['id']));
        } else {
          emit(TokenPurchaseError('Erreur lors de la création du paiement.'));
        }
      } catch (e) {
        emit(TokenPurchaseError(e.toString()));
      }
    });
  }

  Future<void> showPaymentSheet(String clientSecret, Emitter<TokenPurchaseState> emit, String parentId, int tokensAmount, String token) async {
    try {
      // Initialiser le PaymentSheet
      await Stripe.instance.initPaymentSheet(
        paymentSheetParameters: SetupPaymentSheetParameters(
          paymentIntentClientSecret: clientSecret,
        ),
      );

      // Afficher le PaymentSheet
      await Stripe.instance.presentPaymentSheet();

      emit(TokenPurchaseSuccess('Paiement réussi'));

      // Appeler la route pour compléter l'achat
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/api/complete-purchase'),
        headers: {'Authorization': 'Bearer $token', 'Content-Type': 'application/json'},
        body: jsonEncode({
          'amount': tokensAmount * 100, // Convertir en centimes
          'currency': 'eur',
          'parent_id': parentId,
        }),
      );
      } catch (e) {
    // Gérer les erreurs lors de l'affichage du PaymentSheet
    emit(TokenPurchaseError('Erreur lors du traitement du paiement: ${e.toString()}'));
    }
  }

}

