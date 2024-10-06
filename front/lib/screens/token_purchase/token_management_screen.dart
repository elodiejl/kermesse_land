import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/screens/token_purchase/token_purchase_screen.dart';
import 'package:front/services/parent/parent_bloc.dart'; // Importer ton bloc Parent
import 'package:front/services/parent/parent_event.dart';
import 'package:front/services/transaction/transaction_bloc.dart'; // Importer ton bloc Transaction
import 'package:front/services/transaction/transaction_event.dart';
import 'package:jwt_decode/jwt_decode.dart';

import '../../services/parent/parent_state.dart';
import '../../services/transaction/transaction_state.dart'; // Importer les événements de Transaction

class TokenManagementPage extends StatefulWidget {
  final String token;
  const TokenManagementPage({super.key, required this.token});

  @override
  TokenManagementPageState createState() => TokenManagementPageState();

}

class TokenManagementPageState extends State<TokenManagementPage> {
  late int parentId;
  late String parentName;
  late Future<void> fetchParentAndTransactionsFuture;
  bool hasError = false;

  @override
  void initState() {
    super.initState();
    // Récupérer l'ID et le nom du parent depuis le token (ou une autre source)
    parentId = _getParentIdFromToken(widget.token);
    parentName = _getParentNameFromToken(widget.token);

    // Charger les données du parent et les transactions
    fetchParentAndTransactionsFuture = _fetchParentAndTransactions().then((_) {
      context.read<ParentBloc>().add(GetParentById(parentId));
      context.read<TransactionBloc>().add(FetchTransactionsByParentIdEvent(parentId as String));
    }).catchError((error) {
      setState(() {
        hasError = true;
      });
      if (kDebugMode) {
        print('Error fetching parent and transactions: $error');
      }
    });
  }

  int _getParentIdFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    if (kDebugMode) {
      print('Decoded Token: $decodedToken');
    }
    return decodedToken['userId'];  // Assumer que 'userId' est l'ID du parent
  }

  String _getParentNameFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    return decodedToken['username'];  // Assumer que 'username' est le nom du parent
  }

  Future<void> _fetchParentAndTransactions() async {
    // Ici, tu pourrais effectuer des appels au backend pour obtenir plus de données si nécessaire
    return;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Gestion des Tokens'),
        backgroundColor: Colors.cyan,
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // BlocBuilder pour afficher les informations du parent
            BlocBuilder<ParentBloc, ParentState>(
              builder: (context, state) {
                if (state is ParentLoading) {
                  return const Center(child: CircularProgressIndicator());
                } else if (state is ParentSuccess) {
                  final int tokensAvailable = state.parent.tokensAmount;

                  return Container(
                    padding: const EdgeInsets.all(16.0),
                    decoration: BoxDecoration(
                      color: Colors.blue[100],
                      borderRadius: BorderRadius.circular(8.0),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.grey.withOpacity(0.2),
                          spreadRadius: 2,
                          blurRadius: 5,
                          offset: const Offset(0, 3),
                        ),
                      ],
                    ),
                    child: Text(
                      'Jetons disponibles: $tokensAvailable',
                      style: TextStyle(
                        fontSize: 24,
                        fontWeight: FontWeight.bold,
                        color: Colors.blue[800],
                      ),
                    ),
                  );
                } else {
                  return const Text('Erreur lors du chargement des jetons.');
                }
              },
            ),

            const SizedBox(height: 20),

            // BlocBuilder pour afficher la liste des transactions
            Expanded(
              child: BlocBuilder<TransactionBloc, TransactionState>(
                builder: (context, state) {
                  if (state is TransactionLoading) {
                    return const Center(child: CircularProgressIndicator());
                  } else if (state is TransactionsSuccess) {
                    final transactions = state.transactions;

                    return ListView.builder(
                      itemCount: transactions.length,
                      itemBuilder: (context, index) {
                        return Card(
                          margin: const EdgeInsets.symmetric(vertical: 8.0),
                          child: ListTile(
                            title: Text(transactions[index].transactionDate),
                            subtitle: Text(
                                '${transactions[index].tokensAmount} jetons - ${transactions[index].price}€'),
                          ),
                        );
                      },
                    );
                  } else {
                    return const Text('Erreur lors du chargement des transactions.');
                  }
                },
              ),
            ),

            // Bouton pour acheter des jetons
            Padding(
              padding: const EdgeInsets.only(top: 16.0),
              child: ElevatedButton(
                onPressed: () {
                  // Naviguer vers la page d'achat de jetons
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => TokenPurchasePage(token: widget.token),
                    ),
                  );
                },
                style: ElevatedButton.styleFrom(
                  padding: const EdgeInsets.symmetric(vertical: 16.0),
                  textStyle: const TextStyle(fontSize: 18),
                ),
                child: const Text('Acheter des jetons'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

