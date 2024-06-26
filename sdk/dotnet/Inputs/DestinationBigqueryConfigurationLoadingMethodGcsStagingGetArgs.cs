// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationBigqueryConfigurationLoadingMethodGcsStagingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more \n\nhere\n\n.
        /// </summary>
        [Input("credential", required: true)]
        public Input<Inputs.DestinationBigqueryConfigurationLoadingMethodGcsStagingCredentialGetArgs> Credential { get; set; } = null!;

        /// <summary>
        /// The name of the GCS bucket. Read more \n\nhere\n\n.
        /// </summary>
        [Input("gcsBucketName", required: true)]
        public Input<string> GcsBucketName { get; set; } = null!;

        /// <summary>
        /// Directory under the GCS bucket where data will be written.
        /// </summary>
        [Input("gcsBucketPath", required: true)]
        public Input<string> GcsBucketPath { get; set; } = null!;

        /// <summary>
        /// This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly. must be one of ["Delete all tmp files from GCS", "Keep all tmp files in GCS"]; Default: "Delete all tmp files from GCS"
        /// </summary>
        [Input("keepFilesInGcsBucket")]
        public Input<string>? KeepFilesInGcsBucket { get; set; }

        public DestinationBigqueryConfigurationLoadingMethodGcsStagingGetArgs()
        {
        }
        public static new DestinationBigqueryConfigurationLoadingMethodGcsStagingGetArgs Empty => new DestinationBigqueryConfigurationLoadingMethodGcsStagingGetArgs();
    }
}
