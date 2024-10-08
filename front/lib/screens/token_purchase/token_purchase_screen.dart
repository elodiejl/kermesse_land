import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:kermesse_land/services/token_purchase/token_purchase_bloc.dart';
import 'package:flutter_stripe/flutter_stripe.dart';
import 'package:jwt_decode/jwt_decode.dart';

import '../../services/token_purchase/token_purchase_event.dart';
import '../../services/token_purchase/token_purchase_state.dart';

class TokenPurchasePage extends StatefulWidget {
  final String token;

  const TokenPurchasePage({super.key, required this.token});

  @override
  TokenPurchasePageState createState() => TokenPurchasePageState();
}

class TokenPurchasePageState extends State<TokenPurchasePage> {
  int _tokenAmount = 1;
  late String parentId; // Changez-le en String si 'userId' est une chaîne
  bool hasError = false;

  @override
  void initState() {
    super.initState();
    parentId = _getParentIdFromToken(widget.token);
  }

  String _getParentIdFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    if (kDebugMode) {
      print('Decoded Token: $decodedToken');
    }
    return decodedToken['userId'].toString();
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => TokenPurchaseBloc(),
      child: Scaffold(
        appBar: AppBar(title: const Text('Acheter des jetons')),
        body: BlocConsumer<TokenPurchaseBloc, TokenPurchaseState>(
          listener: (context, state) async {
            if (state is TokenPurchaseSuccess) {
              await Stripe.instance.initPaymentSheet(
                paymentSheetParameters: SetupPaymentSheetParameters(
                  paymentIntentClientSecret: state.transactionId,
                  merchantDisplayName: 'Kermesse Land App',
                ),
              );
              await Stripe.instance.presentPaymentSheet();
              ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('Paiement réussi')));
            } else if (state is TokenPurchaseError) {
              ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Erreur : ${state.error}')));
            }
          },
          builder: (context, state) {
            if (state is TokenPurchaseLoading) {
              return const Center(child: CircularProgressIndicator());
            }
            return Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                const Text('Sélectionnez le nombre de jetons'),
                Slider(
                  value: _tokenAmount.toDouble(),
                  min: 1,
                  max: 100,
                  divisions: 100,
                  label: '$_tokenAmount jetons',
                  onChanged: (double value) {
                    setState(() {
                      _tokenAmount = value.toInt();
                    });
                  },
                ),
                ElevatedButton(
                  onPressed: _tokenAmount > 0
                      ? () {
                    BlocProvider.of<TokenPurchaseBloc>(context).add(PurchaseTokens(widget.token,parentId, _tokenAmount));
                  }
                      : null, // Désactive le bouton si le montant de jetons est 0
                  child: Text('Payer $_tokenAmount€'),
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
