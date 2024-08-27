package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "kleo",
		Short: "Kleo is a beeter way to you to take control over your finances.",
		Long: `
        With Kleo:

        * define and monitore your budget
        * track your expenses and incomes
        * find ways to reduces your expenses
        * tacke control over your finances and your life!

        This package is the backen of the app.
        `,
	}
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			// configFilePath, exists := os.LookupEnv("KLEO_CONFIG_PATH")
			// if !exists {
			// 	configFilePath = "config/settings.yaml"
			// }
			// cfg, err := config.NewConfig(configFilePath)
			// if err != nil {
			// 	log.Fatalf("an error occured during configuration loading: %s", err.Error())
			// }
			// adapter := core.InjectGrpc(utils.NewDB(cfg.Database))
			//
			// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
			// if err != nil {
			// 	log.Fatalf("failed to listen: %v", err)
			// }
			// s := grpc.NewServer()
			// generated.RegisterTransactionServiceServer(s, &adapter)
			// log.Printf("server listening at %v", lis.Addr())
			// if err := s.Serve(lis); err != nil {
			// 	log.Fatalf("failed to serve: %v", err)
			// }
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
